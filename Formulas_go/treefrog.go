package main

// Treefrog - High-speed C++ MVC Framework for Web Application
// Homepage: https://www.treefrogframework.org/

import (
	"fmt"
	
	"os/exec"
)

func installTreefrog() {
	// Método 1: Descargar y extraer .tar.gz
	treefrog_tar_url := "https://github.com/treefrogframework/treefrog-framework/archive/refs/tags/v2.9.0.tar.gz"
	treefrog_cmd_tar := exec.Command("curl", "-L", treefrog_tar_url, "-o", "package.tar.gz")
	err := treefrog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	treefrog_zip_url := "https://github.com/treefrogframework/treefrog-framework/archive/refs/tags/v2.9.0.zip"
	treefrog_cmd_zip := exec.Command("curl", "-L", treefrog_zip_url, "-o", "package.zip")
	err = treefrog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	treefrog_bin_url := "https://github.com/treefrogframework/treefrog-framework/archive/refs/tags/v2.9.0.bin"
	treefrog_cmd_bin := exec.Command("curl", "-L", treefrog_bin_url, "-o", "binary.bin")
	err = treefrog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	treefrog_src_url := "https://github.com/treefrogframework/treefrog-framework/archive/refs/tags/v2.9.0.src.tar.gz"
	treefrog_cmd_src := exec.Command("curl", "-L", treefrog_src_url, "-o", "source.tar.gz")
	err = treefrog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	treefrog_cmd_direct := exec.Command("./binary")
	err = treefrog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gflags")
exec.Command("latte", "install", "gflags")
	fmt.Println("Instalando dependencia: glog")
exec.Command("latte", "install", "glog")
	fmt.Println("Instalando dependencia: mongo-c-driver")
exec.Command("latte", "install", "mongo-c-driver")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
}
