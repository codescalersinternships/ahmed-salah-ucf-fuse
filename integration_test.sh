#!/bin/sh

go run ./demo/demo.go mnt &
sleep .5

Name=$(cat ./mnt/Name)
if [[ $Name != 'Salah' ]]; then
    echo 'TEST FAILED: file "Name" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

echo 'TEST #1 PASSED'

find ./mnt/Sub >> /dev/null
if [[ $? != 0 ]]; then
    echo 'TEST FAILED: dir "Sub" does not exist'
fi

echo 'TEST #2 PASSED'

SomeValue=$(cat ./mnt/Sub/SomeValue)
if [[ $SomeValue != 3.14 ]]; then
    echo 'TEST FAILED: file "Float" does not match struct value'
    fusermount -zu ./mnt
    exit 1
fi

echo 'TEST #3 PASSED'

sleep 2
Age=$(cat ./mnt/Age)
if [[ $Age == 22 ]]; then
    echo 'TEST FAILED: file "Age" was not modified'
    fusermount -zu ./mnt
    exit 1
fi

echo 'TEST #4 PASSED'

echo ''

echo 'INTEGRATION TEST PASSED SUCCESSFULLY'