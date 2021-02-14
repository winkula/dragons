Function Dragons-Render($Filename, $Code, $Border) {
    $dragons = "go run ./cmd/dragons"
    $borderArg = if ($Border) {""} else {"-no-outline"}
    Invoke-Expression "$dragons render -filename $filename $borderArg $code"
}

Dragons-Render -Filename "print/instructions/square-dragon" -Code "d" -Border $false
Dragons-Render -Filename "print/instructions/square-fire" -Code "f" -Border $false
Dragons-Render -Filename "print/instructions/square-air" -Code "x" -Border $false
Dragons-Render -Filename "print/instructions/square-nodragon" -Code "n" -Border $false

Dragons-Render -Filename "print/instructions/rule1" -Code "___,_d_,___" -Border $false
Dragons-Render -Filename "print/instructions/rule2" -Code "_____,_df__,__fd_" -Border $false
Dragons-Render -Filename "print/instructions/rule3" -Code "dx_,x__,___" -Border $false

Dragons-Render -Filename "print/instructions/example-step0" -Code "___,_fd,xf_" -Border $true
Dragons-Render -Filename "print/instructions/example-step1" -Code "_nn,_fd,xfn" -Border $true
Dragons-Render -Filename "print/instructions/example-step2" -Code "_nx,_fd,xfx" -Border $true
Dragons-Render -Filename "print/instructions/example-step3" -Code "_nx,dfd,xfx" -Border $true
Dragons-Render -Filename "print/instructions/example-step4" -Code "_fx,dfd,xfx" -Border $true
Dragons-Render -Filename "print/instructions/example-step5" -Code "xfx,dfd,xfx" -Border $true
