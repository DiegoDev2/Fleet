package main

// Uftrace - Function graph tracer for C/C++/Rust
// Homepage: https://uftrace.github.io/slide/

import (
	"fmt"
	
	"os/exec"
)

func installUftrace() {
	// Método 1: Descargar y extraer .tar.gz
	uftrace_tar_url := "https://github.com/namhyung/uftrace/archive/refs/tags/v0.16.tar.gz"
	uftrace_cmd_tar := exec.Command("curl", "-L", uftrace_tar_url, "-o", "package.tar.gz")
	err := uftrace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uftrace_zip_url := "https://github.com/namhyung/uftrace/archive/refs/tags/v0.16.zip"
	uftrace_cmd_zip := exec.Command("curl", "-L", uftrace_zip_url, "-o", "package.zip")
	err = uftrace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uftrace_bin_url := "https://github.com/namhyung/uftrace/archive/refs/tags/v0.16.bin"
	uftrace_cmd_bin := exec.Command("curl", "-L", uftrace_bin_url, "-o", "binary.bin")
	err = uftrace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uftrace_src_url := "https://github.com/namhyung/uftrace/archive/refs/tags/v0.16.src.tar.gz"
	uftrace_cmd_src := exec.Command("curl", "-L", uftrace_src_url, "-o", "source.tar.gz")
	err = uftrace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uftrace_cmd_direct := exec.Command("./binary")
	err = uftrace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: capstone")
	exec.Command("latte", "install", "capstone").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: libunwind")
	exec.Command("latte", "install", "libunwind").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
