package main

// Iproute2 - Linux routing utilities
// Homepage: https://wiki.linuxfoundation.org/networking/iproute2

import (
	"fmt"
	
	"os/exec"
)

func installIproute2() {
	// Método 1: Descargar y extraer .tar.gz
	iproute2_tar_url := "https://mirrors.edge.kernel.org/pub/linux/utils/net/iproute2/iproute2-6.10.0.tar.xz"
	iproute2_cmd_tar := exec.Command("curl", "-L", iproute2_tar_url, "-o", "package.tar.gz")
	err := iproute2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iproute2_zip_url := "https://mirrors.edge.kernel.org/pub/linux/utils/net/iproute2/iproute2-6.10.0.tar.xz"
	iproute2_cmd_zip := exec.Command("curl", "-L", iproute2_zip_url, "-o", "package.zip")
	err = iproute2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iproute2_bin_url := "https://mirrors.edge.kernel.org/pub/linux/utils/net/iproute2/iproute2-6.10.0.tar.xz"
	iproute2_cmd_bin := exec.Command("curl", "-L", iproute2_bin_url, "-o", "binary.bin")
	err = iproute2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iproute2_src_url := "https://mirrors.edge.kernel.org/pub/linux/utils/net/iproute2/iproute2-6.10.0.tar.xz"
	iproute2_cmd_src := exec.Command("curl", "-L", iproute2_src_url, "-o", "source.tar.gz")
	err = iproute2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iproute2_cmd_direct := exec.Command("./binary")
	err = iproute2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: libbpf")
	exec.Command("latte", "install", "libbpf").Run()
	fmt.Println("Instalando dependencia: libcap")
	exec.Command("latte", "install", "libcap").Run()
	fmt.Println("Instalando dependencia: libmnl")
	exec.Command("latte", "install", "libmnl").Run()
	fmt.Println("Instalando dependencia: libtirpc")
	exec.Command("latte", "install", "libtirpc").Run()
}
