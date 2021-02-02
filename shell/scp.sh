#!/usr/bin/expect

set timeout -1
set host [lindex $argv 0]
set username [lindex $argv 1]
set password [lindex $argv 2]
set src_file [lindex $argv 3]
set dest_file [lindex $argv 4]

spawn scp -r $src_file $username@$host:$dest_file
expect {
    "(yes/no)?" {
        send "yes\n"
        exp_continue}
    "password:" {
        send "$password\n"
        exp_continue}
    "Permission denied" {
        send_user "invalid password or account\n"
        exit}
    default {
        send_user "error"
        exit}}
expect eof
