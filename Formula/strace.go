package main

// Strace - Diagnostic, instructional, and debugging tool for the Linux kernel
// Homepage: https://strace.io/

import (
	"fmt"
	
	"os/exec"
)

func installStrace() {
	// Método 1: Descargar y extraer .tar.gz
	strace_tar_url := "https://github.com/strace/strace/releases/download/v6.10/strace-6.10.tar.xz"
	strace_cmd_tar := exec.Command("curl", "-L", strace_tar_url, "-o", "package.tar.gz")
	err := strace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	strace_zip_url := "https://github.com/strace/strace/releases/download/v6.10/strace-6.10.tar.xz"
	strace_cmd_zip := exec.Command("curl", "-L", strace_zip_url, "-o", "package.zip")
	err = strace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	strace_bin_url := "https://github.com/strace/strace/releases/download/v6.10/strace-6.10.tar.xz"
	strace_cmd_bin := exec.Command("curl", "-L", strace_bin_url, "-o", "binary.bin")
	err = strace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	strace_src_url := "https://github.com/strace/strace/releases/download/v6.10/strace-6.10.tar.xz"
	strace_cmd_src := exec.Command("curl", "-L", strace_src_url, "-o", "source.tar.gz")
	err = strace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	strace_cmd_direct := exec.Command("./binary")
	err = strace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: glibc")
	exec.Command("latte", "install", "glibc").Run()
	fmt.Println("Instalando dependencia: linux-headers@5.15")
	exec.Command("latte", "install", "linux-headers@5.15").Run()
}
