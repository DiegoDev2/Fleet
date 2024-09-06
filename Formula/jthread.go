package main

// Jthread - C++ class to make use of threads easy
// Homepage: https://research.edm.uhasselt.be/jori/page/CS/Jthread.html

import (
	"fmt"
	
	"os/exec"
)

func installJthread() {
	// Método 1: Descargar y extraer .tar.gz
	jthread_tar_url := "https://research.edm.uhasselt.be/jori/jthread/jthread-1.3.3.tar.bz2"
	jthread_cmd_tar := exec.Command("curl", "-L", jthread_tar_url, "-o", "package.tar.gz")
	err := jthread_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jthread_zip_url := "https://research.edm.uhasselt.be/jori/jthread/jthread-1.3.3.tar.bz2"
	jthread_cmd_zip := exec.Command("curl", "-L", jthread_zip_url, "-o", "package.zip")
	err = jthread_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jthread_bin_url := "https://research.edm.uhasselt.be/jori/jthread/jthread-1.3.3.tar.bz2"
	jthread_cmd_bin := exec.Command("curl", "-L", jthread_bin_url, "-o", "binary.bin")
	err = jthread_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jthread_src_url := "https://research.edm.uhasselt.be/jori/jthread/jthread-1.3.3.tar.bz2"
	jthread_cmd_src := exec.Command("curl", "-L", jthread_src_url, "-o", "source.tar.gz")
	err = jthread_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jthread_cmd_direct := exec.Command("./binary")
	err = jthread_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
