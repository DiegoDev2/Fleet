package main

// Gptfdisk - Text-mode partitioning tools
// Homepage: https://www.rodsbooks.com/gdisk/

import (
	"fmt"
	
	"os/exec"
)

func installGptfdisk() {
	// Método 1: Descargar y extraer .tar.gz
	gptfdisk_tar_url := "https://downloads.sourceforge.net/project/gptfdisk/gptfdisk/1.0.10/gptfdisk-1.0.10.tar.gz"
	gptfdisk_cmd_tar := exec.Command("curl", "-L", gptfdisk_tar_url, "-o", "package.tar.gz")
	err := gptfdisk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gptfdisk_zip_url := "https://downloads.sourceforge.net/project/gptfdisk/gptfdisk/1.0.10/gptfdisk-1.0.10.zip"
	gptfdisk_cmd_zip := exec.Command("curl", "-L", gptfdisk_zip_url, "-o", "package.zip")
	err = gptfdisk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gptfdisk_bin_url := "https://downloads.sourceforge.net/project/gptfdisk/gptfdisk/1.0.10/gptfdisk-1.0.10.bin"
	gptfdisk_cmd_bin := exec.Command("curl", "-L", gptfdisk_bin_url, "-o", "binary.bin")
	err = gptfdisk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gptfdisk_src_url := "https://downloads.sourceforge.net/project/gptfdisk/gptfdisk/1.0.10/gptfdisk-1.0.10.src.tar.gz"
	gptfdisk_cmd_src := exec.Command("curl", "-L", gptfdisk_src_url, "-o", "source.tar.gz")
	err = gptfdisk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gptfdisk_cmd_direct := exec.Command("./binary")
	err = gptfdisk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: popt")
	exec.Command("latte", "install", "popt").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
