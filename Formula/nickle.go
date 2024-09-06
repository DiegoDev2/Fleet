package main

// Nickle - Desk calculator language
// Homepage: https://www.nickle.org/

import (
	"fmt"
	
	"os/exec"
)

func installNickle() {
	// Método 1: Descargar y extraer .tar.gz
	nickle_tar_url := "https://deb.debian.org/debian/pool/main/n/nickle/nickle_2.97.tar.xz"
	nickle_cmd_tar := exec.Command("curl", "-L", nickle_tar_url, "-o", "package.tar.gz")
	err := nickle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nickle_zip_url := "https://deb.debian.org/debian/pool/main/n/nickle/nickle_2.97.tar.xz"
	nickle_cmd_zip := exec.Command("curl", "-L", nickle_zip_url, "-o", "package.zip")
	err = nickle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nickle_bin_url := "https://deb.debian.org/debian/pool/main/n/nickle/nickle_2.97.tar.xz"
	nickle_cmd_bin := exec.Command("curl", "-L", nickle_bin_url, "-o", "binary.bin")
	err = nickle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nickle_src_url := "https://deb.debian.org/debian/pool/main/n/nickle/nickle_2.97.tar.xz"
	nickle_cmd_src := exec.Command("curl", "-L", nickle_src_url, "-o", "source.tar.gz")
	err = nickle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nickle_cmd_direct := exec.Command("./binary")
	err = nickle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
