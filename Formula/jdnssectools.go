package main

// JdnssecTools - Java command-line tools for DNSSEC
// Homepage: https://github.com/dblacka/jdnssec-tools

import (
	"fmt"
	
	"os/exec"
)

func installJdnssecTools() {
	// Método 1: Descargar y extraer .tar.gz
	jdnssectools_tar_url := "https://github.com/dblacka/jdnssec-tools/releases/download/v0.20/jdnssec-tools-0.20.tar.gz"
	jdnssectools_cmd_tar := exec.Command("curl", "-L", jdnssectools_tar_url, "-o", "package.tar.gz")
	err := jdnssectools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jdnssectools_zip_url := "https://github.com/dblacka/jdnssec-tools/releases/download/v0.20/jdnssec-tools-0.20.zip"
	jdnssectools_cmd_zip := exec.Command("curl", "-L", jdnssectools_zip_url, "-o", "package.zip")
	err = jdnssectools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jdnssectools_bin_url := "https://github.com/dblacka/jdnssec-tools/releases/download/v0.20/jdnssec-tools-0.20.bin"
	jdnssectools_cmd_bin := exec.Command("curl", "-L", jdnssectools_bin_url, "-o", "binary.bin")
	err = jdnssectools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jdnssectools_src_url := "https://github.com/dblacka/jdnssec-tools/releases/download/v0.20/jdnssec-tools-0.20.src.tar.gz"
	jdnssectools_cmd_src := exec.Command("curl", "-L", jdnssectools_src_url, "-o", "source.tar.gz")
	err = jdnssectools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jdnssectools_cmd_direct := exec.Command("./binary")
	err = jdnssectools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
