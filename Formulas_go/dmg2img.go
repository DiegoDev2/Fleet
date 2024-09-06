package main

// Dmg2img - Utilities for converting macOS DMG images
// Homepage: http://vu1tur.eu.org/tools/

import (
	"fmt"
	
	"os/exec"
)

func installDmg2img() {
	// Método 1: Descargar y extraer .tar.gz
	dmg2img_tar_url := "http://vu1tur.eu.org/tools/dmg2img-1.6.7.tar.gz"
	dmg2img_cmd_tar := exec.Command("curl", "-L", dmg2img_tar_url, "-o", "package.tar.gz")
	err := dmg2img_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dmg2img_zip_url := "http://vu1tur.eu.org/tools/dmg2img-1.6.7.zip"
	dmg2img_cmd_zip := exec.Command("curl", "-L", dmg2img_zip_url, "-o", "package.zip")
	err = dmg2img_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dmg2img_bin_url := "http://vu1tur.eu.org/tools/dmg2img-1.6.7.bin"
	dmg2img_cmd_bin := exec.Command("curl", "-L", dmg2img_bin_url, "-o", "binary.bin")
	err = dmg2img_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dmg2img_src_url := "http://vu1tur.eu.org/tools/dmg2img-1.6.7.src.tar.gz"
	dmg2img_cmd_src := exec.Command("curl", "-L", dmg2img_src_url, "-o", "source.tar.gz")
	err = dmg2img_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dmg2img_cmd_direct := exec.Command("./binary")
	err = dmg2img_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
