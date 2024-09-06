package main

// Xmlstarlet - XML command-line utilities
// Homepage: https://xmlstar.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installXmlstarlet() {
	// Método 1: Descargar y extraer .tar.gz
	xmlstarlet_tar_url := "https://downloads.sourceforge.net/project/xmlstar/xmlstarlet/1.6.1/xmlstarlet-1.6.1.tar.gz"
	xmlstarlet_cmd_tar := exec.Command("curl", "-L", xmlstarlet_tar_url, "-o", "package.tar.gz")
	err := xmlstarlet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmlstarlet_zip_url := "https://downloads.sourceforge.net/project/xmlstar/xmlstarlet/1.6.1/xmlstarlet-1.6.1.zip"
	xmlstarlet_cmd_zip := exec.Command("curl", "-L", xmlstarlet_zip_url, "-o", "package.zip")
	err = xmlstarlet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmlstarlet_bin_url := "https://downloads.sourceforge.net/project/xmlstar/xmlstarlet/1.6.1/xmlstarlet-1.6.1.bin"
	xmlstarlet_cmd_bin := exec.Command("curl", "-L", xmlstarlet_bin_url, "-o", "binary.bin")
	err = xmlstarlet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmlstarlet_src_url := "https://downloads.sourceforge.net/project/xmlstar/xmlstarlet/1.6.1/xmlstarlet-1.6.1.src.tar.gz"
	xmlstarlet_cmd_src := exec.Command("curl", "-L", xmlstarlet_src_url, "-o", "source.tar.gz")
	err = xmlstarlet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmlstarlet_cmd_direct := exec.Command("./binary")
	err = xmlstarlet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
