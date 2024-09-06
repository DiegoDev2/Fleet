package main

// Afflib - Advanced Forensic Format
// Homepage: https://github.com/sshock/AFFLIBv3

import (
	"fmt"
	
	"os/exec"
)

func installAfflib() {
	// Método 1: Descargar y extraer .tar.gz
	afflib_tar_url := "https://github.com/sshock/AFFLIBv3/archive/refs/tags/v3.7.20.tar.gz"
	afflib_cmd_tar := exec.Command("curl", "-L", afflib_tar_url, "-o", "package.tar.gz")
	err := afflib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	afflib_zip_url := "https://github.com/sshock/AFFLIBv3/archive/refs/tags/v3.7.20.zip"
	afflib_cmd_zip := exec.Command("curl", "-L", afflib_zip_url, "-o", "package.zip")
	err = afflib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	afflib_bin_url := "https://github.com/sshock/AFFLIBv3/archive/refs/tags/v3.7.20.bin"
	afflib_cmd_bin := exec.Command("curl", "-L", afflib_bin_url, "-o", "binary.bin")
	err = afflib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	afflib_src_url := "https://github.com/sshock/AFFLIBv3/archive/refs/tags/v3.7.20.src.tar.gz"
	afflib_cmd_src := exec.Command("curl", "-L", afflib_src_url, "-o", "source.tar.gz")
	err = afflib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	afflib_cmd_direct := exec.Command("./binary")
	err = afflib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
