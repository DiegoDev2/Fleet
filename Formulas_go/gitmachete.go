package main

// GitMachete - Git repository organizer & rebase workflow automation tool
// Homepage: https://github.com/VirtusLab/git-machete

import (
	"fmt"
	
	"os/exec"
)

func installGitMachete() {
	// Método 1: Descargar y extraer .tar.gz
	gitmachete_tar_url := "https://files.pythonhosted.org/packages/66/e2/f7d9758cd872c5a5d3b49b90f41c1e4b8c8a9db9a3024b9f0c6ea5045e16/git_machete-3.29.2.tar.gz"
	gitmachete_cmd_tar := exec.Command("curl", "-L", gitmachete_tar_url, "-o", "package.tar.gz")
	err := gitmachete_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitmachete_zip_url := "https://files.pythonhosted.org/packages/66/e2/f7d9758cd872c5a5d3b49b90f41c1e4b8c8a9db9a3024b9f0c6ea5045e16/git_machete-3.29.2.zip"
	gitmachete_cmd_zip := exec.Command("curl", "-L", gitmachete_zip_url, "-o", "package.zip")
	err = gitmachete_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitmachete_bin_url := "https://files.pythonhosted.org/packages/66/e2/f7d9758cd872c5a5d3b49b90f41c1e4b8c8a9db9a3024b9f0c6ea5045e16/git_machete-3.29.2.bin"
	gitmachete_cmd_bin := exec.Command("curl", "-L", gitmachete_bin_url, "-o", "binary.bin")
	err = gitmachete_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitmachete_src_url := "https://files.pythonhosted.org/packages/66/e2/f7d9758cd872c5a5d3b49b90f41c1e4b8c8a9db9a3024b9f0c6ea5045e16/git_machete-3.29.2.src.tar.gz"
	gitmachete_cmd_src := exec.Command("curl", "-L", gitmachete_src_url, "-o", "source.tar.gz")
	err = gitmachete_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitmachete_cmd_direct := exec.Command("./binary")
	err = gitmachete_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
