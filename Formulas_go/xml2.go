package main

// Xml2 - Makes XML and HTML more amenable to classic UNIX text tools
// Homepage: https://web.archive.org/web/20160730094113/www.ofb.net/~egnor/xml2/

import (
	"fmt"
	
	"os/exec"
)

func installXml2() {
	// Método 1: Descargar y extraer .tar.gz
	xml2_tar_url := "https://web.archive.org/web/20160427221603/download.ofb.net/gale/xml2-0.5.tar.gz"
	xml2_cmd_tar := exec.Command("curl", "-L", xml2_tar_url, "-o", "package.tar.gz")
	err := xml2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xml2_zip_url := "https://web.archive.org/web/20160427221603/download.ofb.net/gale/xml2-0.5.zip"
	xml2_cmd_zip := exec.Command("curl", "-L", xml2_zip_url, "-o", "package.zip")
	err = xml2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xml2_bin_url := "https://web.archive.org/web/20160427221603/download.ofb.net/gale/xml2-0.5.bin"
	xml2_cmd_bin := exec.Command("curl", "-L", xml2_bin_url, "-o", "binary.bin")
	err = xml2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xml2_src_url := "https://web.archive.org/web/20160427221603/download.ofb.net/gale/xml2-0.5.src.tar.gz"
	xml2_cmd_src := exec.Command("curl", "-L", xml2_src_url, "-o", "source.tar.gz")
	err = xml2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xml2_cmd_direct := exec.Command("./binary")
	err = xml2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
