package main

// GitIf - Glulx interpreter that is optimized for speed
// Homepage: https://ifarchive.org/indexes/if-archiveXprogrammingXglulxXinterpretersXgit.html

import (
	"fmt"
	
	"os/exec"
)

func installGitIf() {
	// Método 1: Descargar y extraer .tar.gz
	gitif_tar_url := "https://ifarchive.org/if-archive/programming/glulx/interpreters/git/git-138.zip"
	gitif_cmd_tar := exec.Command("curl", "-L", gitif_tar_url, "-o", "package.tar.gz")
	err := gitif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitif_zip_url := "https://ifarchive.org/if-archive/programming/glulx/interpreters/git/git-138.zip"
	gitif_cmd_zip := exec.Command("curl", "-L", gitif_zip_url, "-o", "package.zip")
	err = gitif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitif_bin_url := "https://ifarchive.org/if-archive/programming/glulx/interpreters/git/git-138.zip"
	gitif_cmd_bin := exec.Command("curl", "-L", gitif_bin_url, "-o", "binary.bin")
	err = gitif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitif_src_url := "https://ifarchive.org/if-archive/programming/glulx/interpreters/git/git-138.zip"
	gitif_cmd_src := exec.Command("curl", "-L", gitif_src_url, "-o", "source.tar.gz")
	err = gitif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitif_cmd_direct := exec.Command("./binary")
	err = gitif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glktermw")
exec.Command("latte", "install", "glktermw")
}
