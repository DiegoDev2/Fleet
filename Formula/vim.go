package main

// Vim - Vi 'workalike' with many additional features
// Homepage: https://www.vim.org/

import (
	"fmt"
	
	"os/exec"
)

func installVim() {
	// Método 1: Descargar y extraer .tar.gz
	vim_tar_url := "https://github.com/vim/vim/archive/refs/tags/v9.1.0700.tar.gz"
	vim_cmd_tar := exec.Command("curl", "-L", vim_tar_url, "-o", "package.tar.gz")
	err := vim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vim_zip_url := "https://github.com/vim/vim/archive/refs/tags/v9.1.0700.zip"
	vim_cmd_zip := exec.Command("curl", "-L", vim_zip_url, "-o", "package.zip")
	err = vim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vim_bin_url := "https://github.com/vim/vim/archive/refs/tags/v9.1.0700.bin"
	vim_cmd_bin := exec.Command("curl", "-L", vim_bin_url, "-o", "binary.bin")
	err = vim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vim_src_url := "https://github.com/vim/vim/archive/refs/tags/v9.1.0700.src.tar.gz"
	vim_cmd_src := exec.Command("curl", "-L", vim_src_url, "-o", "source.tar.gz")
	err = vim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vim_cmd_direct := exec.Command("./binary")
	err = vim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
	fmt.Println("Instalando dependencia: acl")
	exec.Command("latte", "install", "acl").Run()
}
