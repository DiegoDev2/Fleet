package main

// Icbirc - Proxy IRC client and ICB server
// Homepage: https://www.benzedrine.ch/icbirc.html

import (
	"fmt"
	
	"os/exec"
)

func installIcbirc() {
	// Método 1: Descargar y extraer .tar.gz
	icbirc_tar_url := "https://www.benzedrine.ch/icbirc-2.1.tar.gz"
	icbirc_cmd_tar := exec.Command("curl", "-L", icbirc_tar_url, "-o", "package.tar.gz")
	err := icbirc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	icbirc_zip_url := "https://www.benzedrine.ch/icbirc-2.1.zip"
	icbirc_cmd_zip := exec.Command("curl", "-L", icbirc_zip_url, "-o", "package.zip")
	err = icbirc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	icbirc_bin_url := "https://www.benzedrine.ch/icbirc-2.1.bin"
	icbirc_cmd_bin := exec.Command("curl", "-L", icbirc_bin_url, "-o", "binary.bin")
	err = icbirc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	icbirc_src_url := "https://www.benzedrine.ch/icbirc-2.1.src.tar.gz"
	icbirc_cmd_src := exec.Command("curl", "-L", icbirc_src_url, "-o", "source.tar.gz")
	err = icbirc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	icbirc_cmd_direct := exec.Command("./binary")
	err = icbirc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bsdmake")
	exec.Command("latte", "install", "bsdmake").Run()
}
