package main

// Lifelines - Text-based genealogy software
// Homepage: https://lifelines.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLifelines() {
	// Método 1: Descargar y extraer .tar.gz
	lifelines_tar_url := "https://downloads.sourceforge.net/project/lifelines/lifelines/3.0.62/lifelines-3.0.62.tar.gz"
	lifelines_cmd_tar := exec.Command("curl", "-L", lifelines_tar_url, "-o", "package.tar.gz")
	err := lifelines_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lifelines_zip_url := "https://downloads.sourceforge.net/project/lifelines/lifelines/3.0.62/lifelines-3.0.62.zip"
	lifelines_cmd_zip := exec.Command("curl", "-L", lifelines_zip_url, "-o", "package.zip")
	err = lifelines_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lifelines_bin_url := "https://downloads.sourceforge.net/project/lifelines/lifelines/3.0.62/lifelines-3.0.62.bin"
	lifelines_cmd_bin := exec.Command("curl", "-L", lifelines_bin_url, "-o", "binary.bin")
	err = lifelines_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lifelines_src_url := "https://downloads.sourceforge.net/project/lifelines/lifelines/3.0.62/lifelines-3.0.62.src.tar.gz"
	lifelines_cmd_src := exec.Command("curl", "-L", lifelines_src_url, "-o", "source.tar.gz")
	err = lifelines_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lifelines_cmd_direct := exec.Command("./binary")
	err = lifelines_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
