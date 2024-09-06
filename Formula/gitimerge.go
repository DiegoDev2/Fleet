package main

// GitImerge - Incremental merge for git
// Homepage: https://github.com/mhagger/git-imerge

import (
	"fmt"
	
	"os/exec"
)

func installGitImerge() {
	// Método 1: Descargar y extraer .tar.gz
	gitimerge_tar_url := "https://files.pythonhosted.org/packages/be/f6/ea97fb920d7c3469e4817cfbf9202db98b4a4cdf71d8740e274af57d728c/git-imerge-1.2.0.tar.gz"
	gitimerge_cmd_tar := exec.Command("curl", "-L", gitimerge_tar_url, "-o", "package.tar.gz")
	err := gitimerge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitimerge_zip_url := "https://files.pythonhosted.org/packages/be/f6/ea97fb920d7c3469e4817cfbf9202db98b4a4cdf71d8740e274af57d728c/git-imerge-1.2.0.zip"
	gitimerge_cmd_zip := exec.Command("curl", "-L", gitimerge_zip_url, "-o", "package.zip")
	err = gitimerge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitimerge_bin_url := "https://files.pythonhosted.org/packages/be/f6/ea97fb920d7c3469e4817cfbf9202db98b4a4cdf71d8740e274af57d728c/git-imerge-1.2.0.bin"
	gitimerge_cmd_bin := exec.Command("curl", "-L", gitimerge_bin_url, "-o", "binary.bin")
	err = gitimerge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitimerge_src_url := "https://files.pythonhosted.org/packages/be/f6/ea97fb920d7c3469e4817cfbf9202db98b4a4cdf71d8740e274af57d728c/git-imerge-1.2.0.src.tar.gz"
	gitimerge_cmd_src := exec.Command("curl", "-L", gitimerge_src_url, "-o", "source.tar.gz")
	err = gitimerge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitimerge_cmd_direct := exec.Command("./binary")
	err = gitimerge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
