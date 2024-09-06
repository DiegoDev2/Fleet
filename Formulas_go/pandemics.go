package main

// Pandemics - Converts your markdown document in a simplified framework
// Homepage: https://pandemics.gitlab.io

import (
	"fmt"
	
	"os/exec"
)

func installPandemics() {
	// Método 1: Descargar y extraer .tar.gz
	pandemics_tar_url := "https://registry.npmjs.org/pandemics/-/pandemics-0.12.1.tgz"
	pandemics_cmd_tar := exec.Command("curl", "-L", pandemics_tar_url, "-o", "package.tar.gz")
	err := pandemics_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pandemics_zip_url := "https://registry.npmjs.org/pandemics/-/pandemics-0.12.1.tgz"
	pandemics_cmd_zip := exec.Command("curl", "-L", pandemics_zip_url, "-o", "package.zip")
	err = pandemics_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pandemics_bin_url := "https://registry.npmjs.org/pandemics/-/pandemics-0.12.1.tgz"
	pandemics_cmd_bin := exec.Command("curl", "-L", pandemics_bin_url, "-o", "binary.bin")
	err = pandemics_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pandemics_src_url := "https://registry.npmjs.org/pandemics/-/pandemics-0.12.1.tgz"
	pandemics_cmd_src := exec.Command("curl", "-L", pandemics_src_url, "-o", "source.tar.gz")
	err = pandemics_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pandemics_cmd_direct := exec.Command("./binary")
	err = pandemics_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: librsvg")
exec.Command("latte", "install", "librsvg")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: pandoc-crossref")
exec.Command("latte", "install", "pandoc-crossref")
}
