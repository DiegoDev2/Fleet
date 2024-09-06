package main

// Graphite2 - Smart font renderer for non-Roman scripts
// Homepage: https://graphite.sil.org/

import (
	"fmt"
	
	"os/exec"
)

func installGraphite2() {
	// Método 1: Descargar y extraer .tar.gz
	graphite2_tar_url := "https://github.com/silnrsi/graphite/releases/download/1.3.14/graphite2-1.3.14.tgz"
	graphite2_cmd_tar := exec.Command("curl", "-L", graphite2_tar_url, "-o", "package.tar.gz")
	err := graphite2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	graphite2_zip_url := "https://github.com/silnrsi/graphite/releases/download/1.3.14/graphite2-1.3.14.tgz"
	graphite2_cmd_zip := exec.Command("curl", "-L", graphite2_zip_url, "-o", "package.zip")
	err = graphite2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	graphite2_bin_url := "https://github.com/silnrsi/graphite/releases/download/1.3.14/graphite2-1.3.14.tgz"
	graphite2_cmd_bin := exec.Command("curl", "-L", graphite2_bin_url, "-o", "binary.bin")
	err = graphite2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	graphite2_src_url := "https://github.com/silnrsi/graphite/releases/download/1.3.14/graphite2-1.3.14.tgz"
	graphite2_cmd_src := exec.Command("curl", "-L", graphite2_src_url, "-o", "source.tar.gz")
	err = graphite2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	graphite2_cmd_direct := exec.Command("./binary")
	err = graphite2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
}
