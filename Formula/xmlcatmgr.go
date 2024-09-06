package main

// Xmlcatmgr - Manipulate SGML and XML catalogs
// Homepage: https://xmlcatmgr.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installXmlcatmgr() {
	// Método 1: Descargar y extraer .tar.gz
	xmlcatmgr_tar_url := "https://downloads.sourceforge.net/project/xmlcatmgr/xmlcatmgr/2.2/xmlcatmgr-2.2.tar.gz"
	xmlcatmgr_cmd_tar := exec.Command("curl", "-L", xmlcatmgr_tar_url, "-o", "package.tar.gz")
	err := xmlcatmgr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmlcatmgr_zip_url := "https://downloads.sourceforge.net/project/xmlcatmgr/xmlcatmgr/2.2/xmlcatmgr-2.2.zip"
	xmlcatmgr_cmd_zip := exec.Command("curl", "-L", xmlcatmgr_zip_url, "-o", "package.zip")
	err = xmlcatmgr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmlcatmgr_bin_url := "https://downloads.sourceforge.net/project/xmlcatmgr/xmlcatmgr/2.2/xmlcatmgr-2.2.bin"
	xmlcatmgr_cmd_bin := exec.Command("curl", "-L", xmlcatmgr_bin_url, "-o", "binary.bin")
	err = xmlcatmgr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmlcatmgr_src_url := "https://downloads.sourceforge.net/project/xmlcatmgr/xmlcatmgr/2.2/xmlcatmgr-2.2.src.tar.gz"
	xmlcatmgr_cmd_src := exec.Command("curl", "-L", xmlcatmgr_src_url, "-o", "source.tar.gz")
	err = xmlcatmgr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmlcatmgr_cmd_direct := exec.Command("./binary")
	err = xmlcatmgr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
