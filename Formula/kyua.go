package main

// Kyua - Testing framework for infrastructure software
// Homepage: https://github.com/freebsd/kyua

import (
	"fmt"
	
	"os/exec"
)

func installKyua() {
	// Método 1: Descargar y extraer .tar.gz
	kyua_tar_url := "https://github.com/freebsd/kyua/releases/download/kyua-0.13/kyua-0.13.tar.gz"
	kyua_cmd_tar := exec.Command("curl", "-L", kyua_tar_url, "-o", "package.tar.gz")
	err := kyua_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kyua_zip_url := "https://github.com/freebsd/kyua/releases/download/kyua-0.13/kyua-0.13.zip"
	kyua_cmd_zip := exec.Command("curl", "-L", kyua_zip_url, "-o", "package.zip")
	err = kyua_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kyua_bin_url := "https://github.com/freebsd/kyua/releases/download/kyua-0.13/kyua-0.13.bin"
	kyua_cmd_bin := exec.Command("curl", "-L", kyua_bin_url, "-o", "binary.bin")
	err = kyua_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kyua_src_url := "https://github.com/freebsd/kyua/releases/download/kyua-0.13/kyua-0.13.src.tar.gz"
	kyua_cmd_src := exec.Command("curl", "-L", kyua_src_url, "-o", "source.tar.gz")
	err = kyua_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kyua_cmd_direct := exec.Command("./binary")
	err = kyua_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: atf")
	exec.Command("latte", "install", "atf").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: lutok")
	exec.Command("latte", "install", "lutok").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
