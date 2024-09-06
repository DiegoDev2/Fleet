package main

// Visitors - Web server log analyzer
// Homepage: https://web.archive.org/web/20221105021137/http://www.hping.org/visitors/

import (
	"fmt"
	
	"os/exec"
)

func installVisitors() {
	// Método 1: Descargar y extraer .tar.gz
	visitors_tar_url := "https://web.archive.org/web/20220420184352/http://www.hping.org/visitors/visitors-0.7.tar.gz"
	visitors_cmd_tar := exec.Command("curl", "-L", visitors_tar_url, "-o", "package.tar.gz")
	err := visitors_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	visitors_zip_url := "https://web.archive.org/web/20220420184352/http://www.hping.org/visitors/visitors-0.7.zip"
	visitors_cmd_zip := exec.Command("curl", "-L", visitors_zip_url, "-o", "package.zip")
	err = visitors_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	visitors_bin_url := "https://web.archive.org/web/20220420184352/http://www.hping.org/visitors/visitors-0.7.bin"
	visitors_cmd_bin := exec.Command("curl", "-L", visitors_bin_url, "-o", "binary.bin")
	err = visitors_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	visitors_src_url := "https://web.archive.org/web/20220420184352/http://www.hping.org/visitors/visitors-0.7.src.tar.gz"
	visitors_cmd_src := exec.Command("curl", "-L", visitors_src_url, "-o", "source.tar.gz")
	err = visitors_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	visitors_cmd_direct := exec.Command("./binary")
	err = visitors_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
