package main

// BotanAT2 - Cryptographic algorithms and formats library in C++
// Homepage: https://botan.randombit.net/

import (
	"fmt"
	
	"os/exec"
)

func installBotanAT2() {
	// Método 1: Descargar y extraer .tar.gz
	botanat2_tar_url := "https://botan.randombit.net/releases/Botan-2.19.5.tar.xz"
	botanat2_cmd_tar := exec.Command("curl", "-L", botanat2_tar_url, "-o", "package.tar.gz")
	err := botanat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	botanat2_zip_url := "https://botan.randombit.net/releases/Botan-2.19.5.tar.xz"
	botanat2_cmd_zip := exec.Command("curl", "-L", botanat2_zip_url, "-o", "package.zip")
	err = botanat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	botanat2_bin_url := "https://botan.randombit.net/releases/Botan-2.19.5.tar.xz"
	botanat2_cmd_bin := exec.Command("curl", "-L", botanat2_bin_url, "-o", "binary.bin")
	err = botanat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	botanat2_src_url := "https://botan.randombit.net/releases/Botan-2.19.5.tar.xz"
	botanat2_cmd_src := exec.Command("curl", "-L", botanat2_src_url, "-o", "source.tar.gz")
	err = botanat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	botanat2_cmd_direct := exec.Command("./binary")
	err = botanat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
}
