package main

// Squashfuse - FUSE filesystem to mount squashfs archives
// Homepage: https://github.com/vasi/squashfuse

import (
	"fmt"
	
	"os/exec"
)

func installSquashfuse() {
	// Método 1: Descargar y extraer .tar.gz
	squashfuse_tar_url := "https://github.com/vasi/squashfuse/releases/download/0.5.2/squashfuse-0.5.2.tar.gz"
	squashfuse_cmd_tar := exec.Command("curl", "-L", squashfuse_tar_url, "-o", "package.tar.gz")
	err := squashfuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	squashfuse_zip_url := "https://github.com/vasi/squashfuse/releases/download/0.5.2/squashfuse-0.5.2.zip"
	squashfuse_cmd_zip := exec.Command("curl", "-L", squashfuse_zip_url, "-o", "package.zip")
	err = squashfuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	squashfuse_bin_url := "https://github.com/vasi/squashfuse/releases/download/0.5.2/squashfuse-0.5.2.bin"
	squashfuse_cmd_bin := exec.Command("curl", "-L", squashfuse_bin_url, "-o", "binary.bin")
	err = squashfuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	squashfuse_src_url := "https://github.com/vasi/squashfuse/releases/download/0.5.2/squashfuse-0.5.2.src.tar.gz"
	squashfuse_cmd_src := exec.Command("curl", "-L", squashfuse_src_url, "-o", "source.tar.gz")
	err = squashfuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	squashfuse_cmd_direct := exec.Command("./binary")
	err = squashfuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libfuse")
	exec.Command("latte", "install", "libfuse").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: squashfs")
	exec.Command("latte", "install", "squashfs").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
