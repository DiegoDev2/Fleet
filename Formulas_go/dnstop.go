package main

// Dnstop - Console tool to analyze DNS traffic
// Homepage: http://dns.measurement-factory.com/tools/dnstop/index.html

import (
	"fmt"
	
	"os/exec"
)

func installDnstop() {
	// Método 1: Descargar y extraer .tar.gz
	dnstop_tar_url := "http://dns.measurement-factory.com/tools/dnstop/src/dnstop-20140915.tar.gz"
	dnstop_cmd_tar := exec.Command("curl", "-L", dnstop_tar_url, "-o", "package.tar.gz")
	err := dnstop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnstop_zip_url := "http://dns.measurement-factory.com/tools/dnstop/src/dnstop-20140915.zip"
	dnstop_cmd_zip := exec.Command("curl", "-L", dnstop_zip_url, "-o", "package.zip")
	err = dnstop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnstop_bin_url := "http://dns.measurement-factory.com/tools/dnstop/src/dnstop-20140915.bin"
	dnstop_cmd_bin := exec.Command("curl", "-L", dnstop_bin_url, "-o", "binary.bin")
	err = dnstop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnstop_src_url := "http://dns.measurement-factory.com/tools/dnstop/src/dnstop-20140915.src.tar.gz"
	dnstop_cmd_src := exec.Command("curl", "-L", dnstop_src_url, "-o", "source.tar.gz")
	err = dnstop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnstop_cmd_direct := exec.Command("./binary")
	err = dnstop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
