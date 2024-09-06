package main

// Metaproxy - Z39.50 proxy and router utilizing Yaz toolkit
// Homepage: https://www.indexdata.com/resources/software/metaproxy/

import (
	"fmt"
	
	"os/exec"
)

func installMetaproxy() {
	// Método 1: Descargar y extraer .tar.gz
	metaproxy_tar_url := "https://ftp.indexdata.com/pub/metaproxy/metaproxy-1.21.0.tar.gz"
	metaproxy_cmd_tar := exec.Command("curl", "-L", metaproxy_tar_url, "-o", "package.tar.gz")
	err := metaproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	metaproxy_zip_url := "https://ftp.indexdata.com/pub/metaproxy/metaproxy-1.21.0.zip"
	metaproxy_cmd_zip := exec.Command("curl", "-L", metaproxy_zip_url, "-o", "package.zip")
	err = metaproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	metaproxy_bin_url := "https://ftp.indexdata.com/pub/metaproxy/metaproxy-1.21.0.bin"
	metaproxy_cmd_bin := exec.Command("curl", "-L", metaproxy_bin_url, "-o", "binary.bin")
	err = metaproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	metaproxy_src_url := "https://ftp.indexdata.com/pub/metaproxy/metaproxy-1.21.0.src.tar.gz"
	metaproxy_cmd_src := exec.Command("curl", "-L", metaproxy_src_url, "-o", "source.tar.gz")
	err = metaproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	metaproxy_cmd_direct := exec.Command("./binary")
	err = metaproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: yaz")
	exec.Command("latte", "install", "yaz").Run()
	fmt.Println("Instalando dependencia: yazpp")
	exec.Command("latte", "install", "yazpp").Run()
}
