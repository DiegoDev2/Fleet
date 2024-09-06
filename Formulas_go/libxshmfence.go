package main

// Libxshmfence - X.Org: Shared memory 'SyncFence' synchronization primitive
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxshmfence() {
	// Método 1: Descargar y extraer .tar.gz
	libxshmfence_tar_url := "https://www.x.org/archive/individual/lib/libxshmfence-1.3.2.tar.xz"
	libxshmfence_cmd_tar := exec.Command("curl", "-L", libxshmfence_tar_url, "-o", "package.tar.gz")
	err := libxshmfence_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxshmfence_zip_url := "https://www.x.org/archive/individual/lib/libxshmfence-1.3.2.tar.xz"
	libxshmfence_cmd_zip := exec.Command("curl", "-L", libxshmfence_zip_url, "-o", "package.zip")
	err = libxshmfence_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxshmfence_bin_url := "https://www.x.org/archive/individual/lib/libxshmfence-1.3.2.tar.xz"
	libxshmfence_cmd_bin := exec.Command("curl", "-L", libxshmfence_bin_url, "-o", "binary.bin")
	err = libxshmfence_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxshmfence_src_url := "https://www.x.org/archive/individual/lib/libxshmfence-1.3.2.tar.xz"
	libxshmfence_cmd_src := exec.Command("curl", "-L", libxshmfence_src_url, "-o", "source.tar.gz")
	err = libxshmfence_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxshmfence_cmd_direct := exec.Command("./binary")
	err = libxshmfence_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
