package main

// Pari - Computer algebra system designed for fast computations in number theory
// Homepage: https://pari.math.u-bordeaux.fr/

import (
	"fmt"
	
	"os/exec"
)

func installPari() {
	// Método 1: Descargar y extraer .tar.gz
	pari_tar_url := "https://pari.math.u-bordeaux.fr/pub/pari/unix/pari-2.15.5.tar.gz"
	pari_cmd_tar := exec.Command("curl", "-L", pari_tar_url, "-o", "package.tar.gz")
	err := pari_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pari_zip_url := "https://pari.math.u-bordeaux.fr/pub/pari/unix/pari-2.15.5.zip"
	pari_cmd_zip := exec.Command("curl", "-L", pari_zip_url, "-o", "package.zip")
	err = pari_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pari_bin_url := "https://pari.math.u-bordeaux.fr/pub/pari/unix/pari-2.15.5.bin"
	pari_cmd_bin := exec.Command("curl", "-L", pari_bin_url, "-o", "binary.bin")
	err = pari_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pari_src_url := "https://pari.math.u-bordeaux.fr/pub/pari/unix/pari-2.15.5.src.tar.gz"
	pari_cmd_src := exec.Command("curl", "-L", pari_src_url, "-o", "source.tar.gz")
	err = pari_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pari_cmd_direct := exec.Command("./binary")
	err = pari_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
