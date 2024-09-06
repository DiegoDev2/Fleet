package main

// Negfix8 - Turn scanned negative images into positives
// Homepage: https://web.archive.org/web/20220926032510/https://sites.google.com/site/negfix/

import (
	"fmt"
	
	"os/exec"
)

func installNegfix8() {
	// Método 1: Descargar y extraer .tar.gz
	negfix8_tar_url := "https://web.archive.org/web/20201022025021/https://sites.google.com/site/negfix/downloads/negfix8.3.tgz"
	negfix8_cmd_tar := exec.Command("curl", "-L", negfix8_tar_url, "-o", "package.tar.gz")
	err := negfix8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	negfix8_zip_url := "https://web.archive.org/web/20201022025021/https://sites.google.com/site/negfix/downloads/negfix8.3.tgz"
	negfix8_cmd_zip := exec.Command("curl", "-L", negfix8_zip_url, "-o", "package.zip")
	err = negfix8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	negfix8_bin_url := "https://web.archive.org/web/20201022025021/https://sites.google.com/site/negfix/downloads/negfix8.3.tgz"
	negfix8_cmd_bin := exec.Command("curl", "-L", negfix8_bin_url, "-o", "binary.bin")
	err = negfix8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	negfix8_src_url := "https://web.archive.org/web/20201022025021/https://sites.google.com/site/negfix/downloads/negfix8.3.tgz"
	negfix8_cmd_src := exec.Command("curl", "-L", negfix8_src_url, "-o", "source.tar.gz")
	err = negfix8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	negfix8_cmd_direct := exec.Command("./binary")
	err = negfix8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
}
