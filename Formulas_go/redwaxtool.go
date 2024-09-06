package main

// RedwaxTool - Universal certificate conversion tool
// Homepage: https://redwax.eu/rt/

import (
	"fmt"
	
	"os/exec"
)

func installRedwaxTool() {
	// Método 1: Descargar y extraer .tar.gz
	redwaxtool_tar_url := "https://redwax.eu/dist/rt/redwax-tool-0.9.4.tar.bz2"
	redwaxtool_cmd_tar := exec.Command("curl", "-L", redwaxtool_tar_url, "-o", "package.tar.gz")
	err := redwaxtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redwaxtool_zip_url := "https://redwax.eu/dist/rt/redwax-tool-0.9.4.tar.bz2"
	redwaxtool_cmd_zip := exec.Command("curl", "-L", redwaxtool_zip_url, "-o", "package.zip")
	err = redwaxtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redwaxtool_bin_url := "https://redwax.eu/dist/rt/redwax-tool-0.9.4.tar.bz2"
	redwaxtool_cmd_bin := exec.Command("curl", "-L", redwaxtool_bin_url, "-o", "binary.bin")
	err = redwaxtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redwaxtool_src_url := "https://redwax.eu/dist/rt/redwax-tool-0.9.4.tar.bz2"
	redwaxtool_cmd_src := exec.Command("curl", "-L", redwaxtool_src_url, "-o", "source.tar.gz")
	err = redwaxtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redwaxtool_cmd_direct := exec.Command("./binary")
	err = redwaxtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: apr")
exec.Command("latte", "install", "apr")
	fmt.Println("Instalando dependencia: apr-util")
exec.Command("latte", "install", "apr-util")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
