package main

// Binkd - TCP/IP FTN Mailer
// Homepage: https://2f.ru/binkd/

import (
	"fmt"
	
	"os/exec"
)

func installBinkd() {
	// Método 1: Descargar y extraer .tar.gz
	binkd_tar_url := "https://happy.kiev.ua/pub/fidosoft/mailer/binkd/binkd-1.0.4.tar.gz"
	binkd_cmd_tar := exec.Command("curl", "-L", binkd_tar_url, "-o", "package.tar.gz")
	err := binkd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	binkd_zip_url := "https://happy.kiev.ua/pub/fidosoft/mailer/binkd/binkd-1.0.4.zip"
	binkd_cmd_zip := exec.Command("curl", "-L", binkd_zip_url, "-o", "package.zip")
	err = binkd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	binkd_bin_url := "https://happy.kiev.ua/pub/fidosoft/mailer/binkd/binkd-1.0.4.bin"
	binkd_cmd_bin := exec.Command("curl", "-L", binkd_bin_url, "-o", "binary.bin")
	err = binkd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	binkd_src_url := "https://happy.kiev.ua/pub/fidosoft/mailer/binkd/binkd-1.0.4.src.tar.gz"
	binkd_cmd_src := exec.Command("curl", "-L", binkd_src_url, "-o", "source.tar.gz")
	err = binkd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	binkd_cmd_direct := exec.Command("./binary")
	err = binkd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
