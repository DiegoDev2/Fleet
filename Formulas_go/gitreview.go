package main

// GitReview - Submit git branches to gerrit for review
// Homepage: https://opendev.org/opendev/git-review

import (
	"fmt"
	
	"os/exec"
)

func installGitReview() {
	// Método 1: Descargar y extraer .tar.gz
	gitreview_tar_url := "https://files.pythonhosted.org/packages/79/ae/1c161f8914731ca5a5b3ce0784f5bc47d9a68f4ce33123d431bf30fc90b6/git-review-2.4.0.tar.gz"
	gitreview_cmd_tar := exec.Command("curl", "-L", gitreview_tar_url, "-o", "package.tar.gz")
	err := gitreview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitreview_zip_url := "https://files.pythonhosted.org/packages/79/ae/1c161f8914731ca5a5b3ce0784f5bc47d9a68f4ce33123d431bf30fc90b6/git-review-2.4.0.zip"
	gitreview_cmd_zip := exec.Command("curl", "-L", gitreview_zip_url, "-o", "package.zip")
	err = gitreview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitreview_bin_url := "https://files.pythonhosted.org/packages/79/ae/1c161f8914731ca5a5b3ce0784f5bc47d9a68f4ce33123d431bf30fc90b6/git-review-2.4.0.bin"
	gitreview_cmd_bin := exec.Command("curl", "-L", gitreview_bin_url, "-o", "binary.bin")
	err = gitreview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitreview_src_url := "https://files.pythonhosted.org/packages/79/ae/1c161f8914731ca5a5b3ce0784f5bc47d9a68f4ce33123d431bf30fc90b6/git-review-2.4.0.src.tar.gz"
	gitreview_cmd_src := exec.Command("curl", "-L", gitreview_src_url, "-o", "source.tar.gz")
	err = gitreview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitreview_cmd_direct := exec.Command("./binary")
	err = gitreview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
