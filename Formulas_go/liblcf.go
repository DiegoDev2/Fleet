package main

// Liblcf - Library for RPG Maker 2000/2003 games data
// Homepage: https://easyrpg.org/

import (
	"fmt"
	
	"os/exec"
)

func installLiblcf() {
	// Método 1: Descargar y extraer .tar.gz
	liblcf_tar_url := "https://easyrpg.org/downloads/player/0.8/liblcf-0.8.tar.xz"
	liblcf_cmd_tar := exec.Command("curl", "-L", liblcf_tar_url, "-o", "package.tar.gz")
	err := liblcf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblcf_zip_url := "https://easyrpg.org/downloads/player/0.8/liblcf-0.8.tar.xz"
	liblcf_cmd_zip := exec.Command("curl", "-L", liblcf_zip_url, "-o", "package.zip")
	err = liblcf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblcf_bin_url := "https://easyrpg.org/downloads/player/0.8/liblcf-0.8.tar.xz"
	liblcf_cmd_bin := exec.Command("curl", "-L", liblcf_bin_url, "-o", "binary.bin")
	err = liblcf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblcf_src_url := "https://easyrpg.org/downloads/player/0.8/liblcf-0.8.tar.xz"
	liblcf_cmd_src := exec.Command("curl", "-L", liblcf_src_url, "-o", "source.tar.gz")
	err = liblcf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblcf_cmd_direct := exec.Command("./binary")
	err = liblcf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: expat")
exec.Command("latte", "install", "expat")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
}
