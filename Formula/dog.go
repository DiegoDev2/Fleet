package main

// Dog - Command-line DNS client
// Homepage: https://dns.lookup.dog/

import (
	"fmt"
	
	"os/exec"
)

func installDog() {
	// Método 1: Descargar y extraer .tar.gz
	dog_tar_url := "https://github.com/ogham/dog/archive/refs/tags/v0.1.0.tar.gz"
	dog_cmd_tar := exec.Command("curl", "-L", dog_tar_url, "-o", "package.tar.gz")
	err := dog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dog_zip_url := "https://github.com/ogham/dog/archive/refs/tags/v0.1.0.zip"
	dog_cmd_zip := exec.Command("curl", "-L", dog_zip_url, "-o", "package.zip")
	err = dog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dog_bin_url := "https://github.com/ogham/dog/archive/refs/tags/v0.1.0.bin"
	dog_cmd_bin := exec.Command("curl", "-L", dog_bin_url, "-o", "binary.bin")
	err = dog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dog_src_url := "https://github.com/ogham/dog/archive/refs/tags/v0.1.0.src.tar.gz"
	dog_cmd_src := exec.Command("curl", "-L", dog_src_url, "-o", "source.tar.gz")
	err = dog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dog_cmd_direct := exec.Command("./binary")
	err = dog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: just")
	exec.Command("latte", "install", "just").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@1.1")
	exec.Command("latte", "install", "openssl@1.1").Run()
}
