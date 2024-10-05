BITS 64

; EXAMPLE: nasm -felf64 memexec-nasm.asm &&  ld memexec-nasm.o && cat /bin/id | ./a.out

; memfd_create

mov rax, 319
push 0x00
mov rdi, rsp
xor rsi, rsi

syscall
mov r9, rax

; read

xor rax, rax
sub rsp, 0x400000
xor rdi, rdi
mov rsi, rsp
mov rdx, 0x400000
syscall

; write

mov rdx, rax
mov rdi, r9 ; memfd fd
mov rax, 1
mov rsi, rsp
syscall


; execve

mov rax, 59
lea rdi, [rel path]
xor rsi, rsi
xor rdx, rdx
syscall

path:
 db '/proc/self/fd/3', 0