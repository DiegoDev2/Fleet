package main

// MuRepo - Tool to work with multiple git repositories
// Homepage: https://github.com/fabioz/mu-repo

import (
	"fmt"
	
	"os/exec"
)

func installMuRepo() {
	// Método 1: Descargar y extraer .tar.gz
	murepo_tar_url := "https://files.pythonhosted.org/packages/0d/3d/ddf28cf3beafadb5b3ea45ab882530c1d993b4fc10c0c61d82c8da624f3d/mu_repo-1.9.0.tar.gz"
	murepo_cmd_tar := exec.Command("curl", "-L", murepo_tar_url, "-o", "package.tar.gz")
	err := murepo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	murepo_zip_url := "https://files.pythonhosted.org/packages/0d/3d/ddf28cf3beafadb5b3ea45ab882530c1d993b4fc10c0c61d82c8da624f3d/mu_repo-1.9.0.zip"
	murepo_cmd_zip := exec.Command("curl", "-L", murepo_zip_url, "-o", "package.zip")
	err = murepo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	murepo_bin_url := "https://files.pythonhosted.org/packages/0d/3d/ddf28cf3beafadb5b3ea45ab882530c1d993b4fc10c0c61d82c8da624f3d/mu_repo-1.9.0.bin"
	murepo_cmd_bin := exec.Command("curl", "-L", murepo_bin_url, "-o", "binary.bin")
	err = murepo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	murepo_src_url := "https://files.pythonhosted.org/packages/0d/3d/ddf28cf3beafadb5b3ea45ab882530c1d993b4fc10c0c61d82c8da624f3d/mu_repo-1.9.0.src.tar.gz"
	murepo_cmd_src := exec.Command("curl", "-L", murepo_src_url, "-o", "source.tar.gz")
	err = murepo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	murepo_cmd_direct := exec.Command("./binary")
	err = murepo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
