package main

// Batik - Java-based toolkit for SVG images
// Homepage: https://xmlgraphics.apache.org/batik/

import (
	"fmt"
	
	"os/exec"
)

func installBatik() {
	// Método 1: Descargar y extraer .tar.gz
	batik_tar_url := "https://www.apache.org/dyn/closer.lua?path=xmlgraphics/batik/binaries/batik-bin-1.17.tar.gz"
	batik_cmd_tar := exec.Command("curl", "-L", batik_tar_url, "-o", "package.tar.gz")
	err := batik_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	batik_zip_url := "https://www.apache.org/dyn/closer.lua?path=xmlgraphics/batik/binaries/batik-bin-1.17.zip"
	batik_cmd_zip := exec.Command("curl", "-L", batik_zip_url, "-o", "package.zip")
	err = batik_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	batik_bin_url := "https://www.apache.org/dyn/closer.lua?path=xmlgraphics/batik/binaries/batik-bin-1.17.bin"
	batik_cmd_bin := exec.Command("curl", "-L", batik_bin_url, "-o", "binary.bin")
	err = batik_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	batik_src_url := "https://www.apache.org/dyn/closer.lua?path=xmlgraphics/batik/binaries/batik-bin-1.17.src.tar.gz"
	batik_cmd_src := exec.Command("curl", "-L", batik_src_url, "-o", "source.tar.gz")
	err = batik_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	batik_cmd_direct := exec.Command("./binary")
	err = batik_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
