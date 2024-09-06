package main

// LibpokerEval - C library to evaluate poker hands
// Homepage: https://pokersource.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibpokerEval() {
	// Método 1: Descargar y extraer .tar.gz
	libpokereval_tar_url := "https://launchpad.net/ubuntu/+archive/primary/+sourcefiles/poker-eval/138.0-1/poker-eval_138.0.orig.tar.gz"
	libpokereval_cmd_tar := exec.Command("curl", "-L", libpokereval_tar_url, "-o", "package.tar.gz")
	err := libpokereval_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpokereval_zip_url := "https://launchpad.net/ubuntu/+archive/primary/+sourcefiles/poker-eval/138.0-1/poker-eval_138.0.orig.zip"
	libpokereval_cmd_zip := exec.Command("curl", "-L", libpokereval_zip_url, "-o", "package.zip")
	err = libpokereval_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpokereval_bin_url := "https://launchpad.net/ubuntu/+archive/primary/+sourcefiles/poker-eval/138.0-1/poker-eval_138.0.orig.bin"
	libpokereval_cmd_bin := exec.Command("curl", "-L", libpokereval_bin_url, "-o", "binary.bin")
	err = libpokereval_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpokereval_src_url := "https://launchpad.net/ubuntu/+archive/primary/+sourcefiles/poker-eval/138.0-1/poker-eval_138.0.orig.src.tar.gz"
	libpokereval_cmd_src := exec.Command("curl", "-L", libpokereval_src_url, "-o", "source.tar.gz")
	err = libpokereval_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpokereval_cmd_direct := exec.Command("./binary")
	err = libpokereval_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
