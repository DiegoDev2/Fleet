package main

// LibgrapeLite - C++ library for parallel graph processing
// Homepage: https://github.com/alibaba/libgrape-lite

import (
	"fmt"
	
	"os/exec"
)

func installLibgrapeLite() {
	// Método 1: Descargar y extraer .tar.gz
	libgrapelite_tar_url := "https://github.com/alibaba/libgrape-lite/archive/refs/tags/v0.3.4.tar.gz"
	libgrapelite_cmd_tar := exec.Command("curl", "-L", libgrapelite_tar_url, "-o", "package.tar.gz")
	err := libgrapelite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgrapelite_zip_url := "https://github.com/alibaba/libgrape-lite/archive/refs/tags/v0.3.4.zip"
	libgrapelite_cmd_zip := exec.Command("curl", "-L", libgrapelite_zip_url, "-o", "package.zip")
	err = libgrapelite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgrapelite_bin_url := "https://github.com/alibaba/libgrape-lite/archive/refs/tags/v0.3.4.bin"
	libgrapelite_cmd_bin := exec.Command("curl", "-L", libgrapelite_bin_url, "-o", "binary.bin")
	err = libgrapelite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgrapelite_src_url := "https://github.com/alibaba/libgrape-lite/archive/refs/tags/v0.3.4.src.tar.gz"
	libgrapelite_cmd_src := exec.Command("curl", "-L", libgrapelite_src_url, "-o", "source.tar.gz")
	err = libgrapelite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgrapelite_cmd_direct := exec.Command("./binary")
	err = libgrapelite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: glog")
	exec.Command("latte", "install", "glog").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
}
