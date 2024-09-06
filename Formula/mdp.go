package main

// Mdp - Command-line based markdown presentation tool
// Homepage: https://github.com/visit1985/mdp

import (
	"fmt"
	
	"os/exec"
)

func installMdp() {
	// Método 1: Descargar y extraer .tar.gz
	mdp_tar_url := "https://github.com/visit1985/mdp/archive/refs/tags/1.0.15.tar.gz"
	mdp_cmd_tar := exec.Command("curl", "-L", mdp_tar_url, "-o", "package.tar.gz")
	err := mdp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdp_zip_url := "https://github.com/visit1985/mdp/archive/refs/tags/1.0.15.zip"
	mdp_cmd_zip := exec.Command("curl", "-L", mdp_zip_url, "-o", "package.zip")
	err = mdp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdp_bin_url := "https://github.com/visit1985/mdp/archive/refs/tags/1.0.15.bin"
	mdp_cmd_bin := exec.Command("curl", "-L", mdp_bin_url, "-o", "binary.bin")
	err = mdp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdp_src_url := "https://github.com/visit1985/mdp/archive/refs/tags/1.0.15.src.tar.gz"
	mdp_cmd_src := exec.Command("curl", "-L", mdp_src_url, "-o", "source.tar.gz")
	err = mdp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdp_cmd_direct := exec.Command("./binary")
	err = mdp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
