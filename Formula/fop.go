package main

// Fop - XSL-FO print formatter for making PDF or PS documents
// Homepage: https://xmlgraphics.apache.org/fop/index.html

import (
	"fmt"
	
	"os/exec"
)

func installFop() {
	// Método 1: Descargar y extraer .tar.gz
	fop_tar_url := "https://www.apache.org/dyn/closer.lua?path=xmlgraphics/fop/binaries/fop-2.9-bin.tar.gz"
	fop_cmd_tar := exec.Command("curl", "-L", fop_tar_url, "-o", "package.tar.gz")
	err := fop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fop_zip_url := "https://www.apache.org/dyn/closer.lua?path=xmlgraphics/fop/binaries/fop-2.9-bin.zip"
	fop_cmd_zip := exec.Command("curl", "-L", fop_zip_url, "-o", "package.zip")
	err = fop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fop_bin_url := "https://www.apache.org/dyn/closer.lua?path=xmlgraphics/fop/binaries/fop-2.9-bin.bin"
	fop_cmd_bin := exec.Command("curl", "-L", fop_bin_url, "-o", "binary.bin")
	err = fop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fop_src_url := "https://www.apache.org/dyn/closer.lua?path=xmlgraphics/fop/binaries/fop-2.9-bin.src.tar.gz"
	fop_cmd_src := exec.Command("curl", "-L", fop_src_url, "-o", "source.tar.gz")
	err = fop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fop_cmd_direct := exec.Command("./binary")
	err = fop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
