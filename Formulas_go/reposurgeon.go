package main

// Reposurgeon - Edit version-control repository history
// Homepage: http://www.catb.org/esr/reposurgeon/

import (
	"fmt"
	
	"os/exec"
)

func installReposurgeon() {
	// Método 1: Descargar y extraer .tar.gz
	reposurgeon_tar_url := "https://gitlab.com/esr/reposurgeon/-/archive/5.1/reposurgeon-5.1.tar.gz"
	reposurgeon_cmd_tar := exec.Command("curl", "-L", reposurgeon_tar_url, "-o", "package.tar.gz")
	err := reposurgeon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reposurgeon_zip_url := "https://gitlab.com/esr/reposurgeon/-/archive/5.1/reposurgeon-5.1.zip"
	reposurgeon_cmd_zip := exec.Command("curl", "-L", reposurgeon_zip_url, "-o", "package.zip")
	err = reposurgeon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reposurgeon_bin_url := "https://gitlab.com/esr/reposurgeon/-/archive/5.1/reposurgeon-5.1.bin"
	reposurgeon_cmd_bin := exec.Command("curl", "-L", reposurgeon_bin_url, "-o", "binary.bin")
	err = reposurgeon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reposurgeon_src_url := "https://gitlab.com/esr/reposurgeon/-/archive/5.1/reposurgeon-5.1.src.tar.gz"
	reposurgeon_cmd_src := exec.Command("curl", "-L", reposurgeon_src_url, "-o", "source.tar.gz")
	err = reposurgeon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reposurgeon_cmd_direct := exec.Command("./binary")
	err = reposurgeon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
exec.Command("latte", "install", "asciidoctor")
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: git")
exec.Command("latte", "install", "git")
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
}
