package main

// Whatmp3 - Small script to create mp3 torrents out of FLACs
// Homepage: https://github.com/RecursiveForest/whatmp3

import (
	"fmt"
	
	"os/exec"
)

func installWhatmp3() {
	// Método 1: Descargar y extraer .tar.gz
	whatmp3_tar_url := "https://github.com/RecursiveForest/whatmp3/archive/refs/tags/v3.8.tar.gz"
	whatmp3_cmd_tar := exec.Command("curl", "-L", whatmp3_tar_url, "-o", "package.tar.gz")
	err := whatmp3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	whatmp3_zip_url := "https://github.com/RecursiveForest/whatmp3/archive/refs/tags/v3.8.zip"
	whatmp3_cmd_zip := exec.Command("curl", "-L", whatmp3_zip_url, "-o", "package.zip")
	err = whatmp3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	whatmp3_bin_url := "https://github.com/RecursiveForest/whatmp3/archive/refs/tags/v3.8.bin"
	whatmp3_cmd_bin := exec.Command("curl", "-L", whatmp3_bin_url, "-o", "binary.bin")
	err = whatmp3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	whatmp3_src_url := "https://github.com/RecursiveForest/whatmp3/archive/refs/tags/v3.8.src.tar.gz"
	whatmp3_cmd_src := exec.Command("curl", "-L", whatmp3_src_url, "-o", "source.tar.gz")
	err = whatmp3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	whatmp3_cmd_direct := exec.Command("./binary")
	err = whatmp3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: mktorrent")
exec.Command("latte", "install", "mktorrent")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
