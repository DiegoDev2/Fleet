package main

// GitOctopus - Continuous merge workflow
// Homepage: https://github.com/lesfurets/git-octopus

import (
	"fmt"
	
	"os/exec"
)

func installGitOctopus() {
	// Método 1: Descargar y extraer .tar.gz
	gitoctopus_tar_url := "https://github.com/lesfurets/git-octopus/archive/refs/tags/v1.4.tar.gz"
	gitoctopus_cmd_tar := exec.Command("curl", "-L", gitoctopus_tar_url, "-o", "package.tar.gz")
	err := gitoctopus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitoctopus_zip_url := "https://github.com/lesfurets/git-octopus/archive/refs/tags/v1.4.zip"
	gitoctopus_cmd_zip := exec.Command("curl", "-L", gitoctopus_zip_url, "-o", "package.zip")
	err = gitoctopus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitoctopus_bin_url := "https://github.com/lesfurets/git-octopus/archive/refs/tags/v1.4.bin"
	gitoctopus_cmd_bin := exec.Command("curl", "-L", gitoctopus_bin_url, "-o", "binary.bin")
	err = gitoctopus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitoctopus_src_url := "https://github.com/lesfurets/git-octopus/archive/refs/tags/v1.4.src.tar.gz"
	gitoctopus_cmd_src := exec.Command("curl", "-L", gitoctopus_src_url, "-o", "source.tar.gz")
	err = gitoctopus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitoctopus_cmd_direct := exec.Command("./binary")
	err = gitoctopus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
