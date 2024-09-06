package main

// Bgpstream - For live and historical BGP data analysis
// Homepage: https://bgpstream.caida.org/

import (
	"fmt"
	
	"os/exec"
)

func installBgpstream() {
	// Método 1: Descargar y extraer .tar.gz
	bgpstream_tar_url := "https://github.com/CAIDA/libbgpstream/releases/download/v2.3.0/libbgpstream-2.3.0.tar.gz"
	bgpstream_cmd_tar := exec.Command("curl", "-L", bgpstream_tar_url, "-o", "package.tar.gz")
	err := bgpstream_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bgpstream_zip_url := "https://github.com/CAIDA/libbgpstream/releases/download/v2.3.0/libbgpstream-2.3.0.zip"
	bgpstream_cmd_zip := exec.Command("curl", "-L", bgpstream_zip_url, "-o", "package.zip")
	err = bgpstream_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bgpstream_bin_url := "https://github.com/CAIDA/libbgpstream/releases/download/v2.3.0/libbgpstream-2.3.0.bin"
	bgpstream_cmd_bin := exec.Command("curl", "-L", bgpstream_bin_url, "-o", "binary.bin")
	err = bgpstream_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bgpstream_src_url := "https://github.com/CAIDA/libbgpstream/releases/download/v2.3.0/libbgpstream-2.3.0.src.tar.gz"
	bgpstream_cmd_src := exec.Command("curl", "-L", bgpstream_src_url, "-o", "source.tar.gz")
	err = bgpstream_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bgpstream_cmd_direct := exec.Command("./binary")
	err = bgpstream_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: librdkafka")
	exec.Command("latte", "install", "librdkafka").Run()
	fmt.Println("Instalando dependencia: wandio")
	exec.Command("latte", "install", "wandio").Run()
}
