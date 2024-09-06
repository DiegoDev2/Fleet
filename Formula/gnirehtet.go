package main

// Gnirehtet - Reverse tethering tool for Android
// Homepage: https://github.com/Genymobile/gnirehtet

import (
	"fmt"
	
	"os/exec"
)

func installGnirehtet() {
	// Método 1: Descargar y extraer .tar.gz
	gnirehtet_tar_url := "https://github.com/Genymobile/gnirehtet/archive/refs/tags/v2.5.1.tar.gz"
	gnirehtet_cmd_tar := exec.Command("curl", "-L", gnirehtet_tar_url, "-o", "package.tar.gz")
	err := gnirehtet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnirehtet_zip_url := "https://github.com/Genymobile/gnirehtet/archive/refs/tags/v2.5.1.zip"
	gnirehtet_cmd_zip := exec.Command("curl", "-L", gnirehtet_zip_url, "-o", "package.zip")
	err = gnirehtet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnirehtet_bin_url := "https://github.com/Genymobile/gnirehtet/archive/refs/tags/v2.5.1.bin"
	gnirehtet_cmd_bin := exec.Command("curl", "-L", gnirehtet_bin_url, "-o", "binary.bin")
	err = gnirehtet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnirehtet_src_url := "https://github.com/Genymobile/gnirehtet/archive/refs/tags/v2.5.1.src.tar.gz"
	gnirehtet_cmd_src := exec.Command("curl", "-L", gnirehtet_src_url, "-o", "source.tar.gz")
	err = gnirehtet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnirehtet_cmd_direct := exec.Command("./binary")
	err = gnirehtet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: socat")
	exec.Command("latte", "install", "socat").Run()
}
