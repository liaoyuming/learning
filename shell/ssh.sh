#!/usr/bin/expect

set timeout 30
set host [lindex $argv 0]
set username [lindex $argv 1]
set password [lindex $argv 2]
set command [lindex $argv 3]

spawn ssh $username@$host  "/root/deploy.sh"
expect {
    "(yes/no)?" {
        send "yes\n"
        exp_continue}
    password: {
        send "$password\n"
        exp_continue}
    "Permission denied" {
        send_user "invalid password or account\n"
        exit}
    timeout {
        send_user "timeout\n"
        exit}
    default {
        send_user "error"
        exit}}
expect eof
