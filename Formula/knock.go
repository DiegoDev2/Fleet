package main

// Knock - Port-knock server
// Homepage: https://github.com/jvinet/knock

import (
	"fmt"
	
	"os/exec"
)

func installKnock() {
	// Método 1: Descargar y extraer .tar.gz
	knock_tar_url := "https://github.com/jvinet/knock/releases/download/v0.8/knock-0.8.tar.gz"
	knock_cmd_tar := exec.Command("curl", "-L", knock_tar_url, "-o", "package.tar.gz")
	err := knock_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	knock_zip_url := "https://github.com/jvinet/knock/releases/download/v0.8/knock-0.8.zip"
	knock_cmd_zip := exec.Command("curl", "-L", knock_zip_url, "-o", "package.zip")
	err = knock_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	knock_bin_url := "https://github.com/jvinet/knock/releases/download/v0.8/knock-0.8.bin"
	knock_cmd_bin := exec.Command("curl", "-L", knock_bin_url, "-o", "binary.bin")
	err = knock_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	knock_src_url := "https://github.com/jvinet/knock/releases/download/v0.8/knock-0.8.src.tar.gz"
	knock_cmd_src := exec.Command("curl", "-L", knock_src_url, "-o", "source.tar.gz")
	err = knock_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	knock_cmd_direct := exec.Command("./binary")
	err = knock_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
