package main

// Kakoune - Selection-based modal text editor
// Homepage: https://github.com/mawww/kakoune

import (
	"fmt"
	
	"os/exec"
)

func installKakoune() {
	// Método 1: Descargar y extraer .tar.gz
	kakoune_tar_url := "https://github.com/mawww/kakoune/releases/download/v2024.05.18/kakoune-2024.05.18.tar.bz2"
	kakoune_cmd_tar := exec.Command("curl", "-L", kakoune_tar_url, "-o", "package.tar.gz")
	err := kakoune_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kakoune_zip_url := "https://github.com/mawww/kakoune/releases/download/v2024.05.18/kakoune-2024.05.18.tar.bz2"
	kakoune_cmd_zip := exec.Command("curl", "-L", kakoune_zip_url, "-o", "package.zip")
	err = kakoune_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kakoune_bin_url := "https://github.com/mawww/kakoune/releases/download/v2024.05.18/kakoune-2024.05.18.tar.bz2"
	kakoune_cmd_bin := exec.Command("curl", "-L", kakoune_bin_url, "-o", "binary.bin")
	err = kakoune_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kakoune_src_url := "https://github.com/mawww/kakoune/releases/download/v2024.05.18/kakoune-2024.05.18.tar.bz2"
	kakoune_cmd_src := exec.Command("curl", "-L", kakoune_src_url, "-o", "source.tar.gz")
	err = kakoune_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kakoune_cmd_direct := exec.Command("./binary")
	err = kakoune_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: binutils")
exec.Command("latte", "install", "binutils")
}
