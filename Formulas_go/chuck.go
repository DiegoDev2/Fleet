package main

// Chuck - Concurrent, on-the-fly audio programming language
// Homepage: https://chuck.cs.princeton.edu/

import (
	"fmt"
	
	"os/exec"
)

func installChuck() {
	// Método 1: Descargar y extraer .tar.gz
	chuck_tar_url := "https://chuck.cs.princeton.edu/release/files/chuck-1.5.2.5.tgz"
	chuck_cmd_tar := exec.Command("curl", "-L", chuck_tar_url, "-o", "package.tar.gz")
	err := chuck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chuck_zip_url := "https://chuck.cs.princeton.edu/release/files/chuck-1.5.2.5.tgz"
	chuck_cmd_zip := exec.Command("curl", "-L", chuck_zip_url, "-o", "package.zip")
	err = chuck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chuck_bin_url := "https://chuck.cs.princeton.edu/release/files/chuck-1.5.2.5.tgz"
	chuck_cmd_bin := exec.Command("curl", "-L", chuck_bin_url, "-o", "binary.bin")
	err = chuck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chuck_src_url := "https://chuck.cs.princeton.edu/release/files/chuck-1.5.2.5.tgz"
	chuck_cmd_src := exec.Command("curl", "-L", chuck_src_url, "-o", "source.tar.gz")
	err = chuck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chuck_cmd_direct := exec.Command("./binary")
	err = chuck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
}
