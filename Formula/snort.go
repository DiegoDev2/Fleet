package main

// Snort - Flexible Network Intrusion Detection System
// Homepage: https://www.snort.org

import (
	"fmt"
	
	"os/exec"
)

func installSnort() {
	// Método 1: Descargar y extraer .tar.gz
	snort_tar_url := "https://github.com/snort3/snort3/archive/refs/tags/3.3.4.0.tar.gz"
	snort_cmd_tar := exec.Command("curl", "-L", snort_tar_url, "-o", "package.tar.gz")
	err := snort_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snort_zip_url := "https://github.com/snort3/snort3/archive/refs/tags/3.3.4.0.zip"
	snort_cmd_zip := exec.Command("curl", "-L", snort_zip_url, "-o", "package.zip")
	err = snort_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snort_bin_url := "https://github.com/snort3/snort3/archive/refs/tags/3.3.4.0.bin"
	snort_cmd_bin := exec.Command("curl", "-L", snort_bin_url, "-o", "binary.bin")
	err = snort_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snort_src_url := "https://github.com/snort3/snort3/archive/refs/tags/3.3.4.0.src.tar.gz"
	snort_cmd_src := exec.Command("curl", "-L", snort_src_url, "-o", "source.tar.gz")
	err = snort_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snort_cmd_direct := exec.Command("./binary")
	err = snort_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: daq")
	exec.Command("latte", "install", "daq").Run()
	fmt.Println("Instalando dependencia: hwloc")
	exec.Command("latte", "install", "hwloc").Run()
	fmt.Println("Instalando dependencia: jemalloc")
	exec.Command("latte", "install", "jemalloc").Run()
	fmt.Println("Instalando dependencia: libdnet")
	exec.Command("latte", "install", "libdnet").Run()
	fmt.Println("Instalando dependencia: libpcap")
	exec.Command("latte", "install", "libpcap").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: vectorscan")
	exec.Command("latte", "install", "vectorscan").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: libunwind")
	exec.Command("latte", "install", "libunwind").Run()
}
