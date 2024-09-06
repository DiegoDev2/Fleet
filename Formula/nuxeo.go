package main

// Nuxeo - Enterprise Content Management
// Homepage: https://nuxeo.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installNuxeo() {
	// Método 1: Descargar y extraer .tar.gz
	nuxeo_tar_url := "https://cdn.nuxeo.com/nuxeo-10.10/nuxeo-server-10.10-tomcat.zip"
	nuxeo_cmd_tar := exec.Command("curl", "-L", nuxeo_tar_url, "-o", "package.tar.gz")
	err := nuxeo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nuxeo_zip_url := "https://cdn.nuxeo.com/nuxeo-10.10/nuxeo-server-10.10-tomcat.zip"
	nuxeo_cmd_zip := exec.Command("curl", "-L", nuxeo_zip_url, "-o", "package.zip")
	err = nuxeo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nuxeo_bin_url := "https://cdn.nuxeo.com/nuxeo-10.10/nuxeo-server-10.10-tomcat.zip"
	nuxeo_cmd_bin := exec.Command("curl", "-L", nuxeo_bin_url, "-o", "binary.bin")
	err = nuxeo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nuxeo_src_url := "https://cdn.nuxeo.com/nuxeo-10.10/nuxeo-server-10.10-tomcat.zip"
	nuxeo_cmd_src := exec.Command("curl", "-L", nuxeo_src_url, "-o", "source.tar.gz")
	err = nuxeo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nuxeo_cmd_direct := exec.Command("./binary")
	err = nuxeo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: exiftool")
	exec.Command("latte", "install", "exiftool").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
	fmt.Println("Instalando dependencia: libwpd")
	exec.Command("latte", "install", "libwpd").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
}
