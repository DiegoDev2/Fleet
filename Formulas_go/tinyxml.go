package main

// Tinyxml - XML parser
// Homepage: http://www.grinninglizard.com/tinyxml/

import (
	"fmt"
	
	"os/exec"
)

func installTinyxml() {
	// Método 1: Descargar y extraer .tar.gz
	tinyxml_tar_url := "https://downloads.sourceforge.net/project/tinyxml/tinyxml/2.6.2/tinyxml_2_6_2.tar.gz"
	tinyxml_cmd_tar := exec.Command("curl", "-L", tinyxml_tar_url, "-o", "package.tar.gz")
	err := tinyxml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tinyxml_zip_url := "https://downloads.sourceforge.net/project/tinyxml/tinyxml/2.6.2/tinyxml_2_6_2.zip"
	tinyxml_cmd_zip := exec.Command("curl", "-L", tinyxml_zip_url, "-o", "package.zip")
	err = tinyxml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tinyxml_bin_url := "https://downloads.sourceforge.net/project/tinyxml/tinyxml/2.6.2/tinyxml_2_6_2.bin"
	tinyxml_cmd_bin := exec.Command("curl", "-L", tinyxml_bin_url, "-o", "binary.bin")
	err = tinyxml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tinyxml_src_url := "https://downloads.sourceforge.net/project/tinyxml/tinyxml/2.6.2/tinyxml_2_6_2.src.tar.gz"
	tinyxml_cmd_src := exec.Command("curl", "-L", tinyxml_src_url, "-o", "source.tar.gz")
	err = tinyxml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tinyxml_cmd_direct := exec.Command("./binary")
	err = tinyxml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
