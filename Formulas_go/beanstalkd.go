package main

// Beanstalkd - Generic work queue originally designed to reduce web latency
// Homepage: https://beanstalkd.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installBeanstalkd() {
	// Método 1: Descargar y extraer .tar.gz
	beanstalkd_tar_url := "https://github.com/beanstalkd/beanstalkd/archive/refs/tags/v1.13.tar.gz"
	beanstalkd_cmd_tar := exec.Command("curl", "-L", beanstalkd_tar_url, "-o", "package.tar.gz")
	err := beanstalkd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	beanstalkd_zip_url := "https://github.com/beanstalkd/beanstalkd/archive/refs/tags/v1.13.zip"
	beanstalkd_cmd_zip := exec.Command("curl", "-L", beanstalkd_zip_url, "-o", "package.zip")
	err = beanstalkd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	beanstalkd_bin_url := "https://github.com/beanstalkd/beanstalkd/archive/refs/tags/v1.13.bin"
	beanstalkd_cmd_bin := exec.Command("curl", "-L", beanstalkd_bin_url, "-o", "binary.bin")
	err = beanstalkd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	beanstalkd_src_url := "https://github.com/beanstalkd/beanstalkd/archive/refs/tags/v1.13.src.tar.gz"
	beanstalkd_cmd_src := exec.Command("curl", "-L", beanstalkd_src_url, "-o", "source.tar.gz")
	err = beanstalkd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	beanstalkd_cmd_direct := exec.Command("./binary")
	err = beanstalkd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
