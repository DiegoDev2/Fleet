package main

// Proteinortho - Detecting orthologous genes within different species
// Homepage: https://gitlab.com/paulklemm_PHD/proteinortho

import (
	"fmt"
	
	"os/exec"
)

func installProteinortho() {
	// Método 1: Descargar y extraer .tar.gz
	proteinortho_tar_url := "https://gitlab.com/paulklemm_PHD/proteinortho/-/archive/v6.3.2/proteinortho-v6.3.2.tar.gz"
	proteinortho_cmd_tar := exec.Command("curl", "-L", proteinortho_tar_url, "-o", "package.tar.gz")
	err := proteinortho_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proteinortho_zip_url := "https://gitlab.com/paulklemm_PHD/proteinortho/-/archive/v6.3.2/proteinortho-v6.3.2.zip"
	proteinortho_cmd_zip := exec.Command("curl", "-L", proteinortho_zip_url, "-o", "package.zip")
	err = proteinortho_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proteinortho_bin_url := "https://gitlab.com/paulklemm_PHD/proteinortho/-/archive/v6.3.2/proteinortho-v6.3.2.bin"
	proteinortho_cmd_bin := exec.Command("curl", "-L", proteinortho_bin_url, "-o", "binary.bin")
	err = proteinortho_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proteinortho_src_url := "https://gitlab.com/paulklemm_PHD/proteinortho/-/archive/v6.3.2/proteinortho-v6.3.2.src.tar.gz"
	proteinortho_cmd_src := exec.Command("curl", "-L", proteinortho_src_url, "-o", "source.tar.gz")
	err = proteinortho_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proteinortho_cmd_direct := exec.Command("./binary")
	err = proteinortho_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: diamond")
	exec.Command("latte", "install", "diamond").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
